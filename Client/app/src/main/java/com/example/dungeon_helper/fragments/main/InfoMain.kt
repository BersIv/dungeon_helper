package com.example.dungeon_helper.fragments.main

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.text.Spannable
import android.text.SpannableString
import android.text.style.ForegroundColorSpan
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import android.widget.SearchView
import com.example.dungeon_helper.R
import androidx.recyclerview.widget.RecyclerView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.databinding.FragmentInfoMainBinding

class InfoMain : Fragment() {

    companion object {
        fun newInstance() = InfoMain()
    }
    private lateinit var viewModel: InfoMainViewModel


    private var _binding: FragmentInfoMainBinding? = null
    private  val binding get() = _binding!!
    private lateinit var gameTitleView: TextView

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val infoMainViewModel = ViewModelProvider(this)[InfoMainViewModel::class.java]

        _binding = FragmentInfoMainBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textInfo
        infoMainViewModel.text.observe(viewLifecycleOwner) {
            textView.text = it
        }


        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        gameTitleView = binding.gameTitle
        gameTitleView.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_navigation_info_to_infoVers)
        }
        val searchView = binding.searchView
        searchView.setOnQueryTextListener(object : SearchView.OnQueryTextListener {
            override fun onQueryTextSubmit(query: String?): Boolean {

                return false
            }

            override fun onQueryTextChange(newText: String?): Boolean {

                if (newText.isNullOrEmpty()) {
                    return true
                }

                val gameTitleText = gameTitleView.text.toString()
                val spannableString = SpannableString(gameTitleText)


                val startIndex = gameTitleText.indexOf(newText, ignoreCase = true)
                val endIndex = startIndex + newText.length

                if (startIndex != -1) {
                    spannableString.setSpan(
                        ForegroundColorSpan(resources.getColor(R.color.purple_700)),
                        startIndex,
                        endIndex,
                        Spannable.SPAN_EXCLUSIVE_EXCLUSIVE
                    )
                }


                gameTitleView.text = spannableString

                return true
            }
        })

    }


}