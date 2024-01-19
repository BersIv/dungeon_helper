package com.example.dungeon_helper.fragments.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import android.text.Spannable
import android.text.SpannableString
import android.text.style.ForegroundColorSpan
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.SearchView
import com.example.dungeon_helper.MasterLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentMasterLobbyInfoContentBinding

class MasterLobbyInfoContent : Fragment() {

    companion object {
        fun newInstance() = MasterLobbyInfoContent()
    }

    private lateinit var viewModel: MasterLobbyInfoContentViewModel
    private var _binding: FragmentMasterLobbyInfoContentBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
      val masterLobbyInfoContentViewModel = ViewModelProvider(this)[MasterLobbyInfoContentViewModel::class.java]
        _binding = FragmentMasterLobbyInfoContentBinding.inflate(inflater, container,false)
        val root: View = binding.root
        return root
    }

    override fun onDestroyView() {
        super.onDestroyView()
        _binding = null
    }

    override fun onStart() {
        super.onStart()
        val backBtn = binding.backBtn
        backBtn.setOnClickListener {
            (activity as MasterLobbyActivity).navController.navigate(R.id.action_masterLobbyInfoContent_to_masterLobbyInfoAllChapters)
        }

        val chapterContent = binding.chapterContent
        val chapterSearchView = binding.chapterSearchView
        chapterSearchView.setOnQueryTextListener(object : SearchView.OnQueryTextListener {
            override fun onQueryTextSubmit(query: String?): Boolean {

                return false
            }

            override fun onQueryTextChange(newText: String?): Boolean {
                val contentText = chapterContent.text.toString()


                chapterContent.text = SpannableString.valueOf(contentText)


                if (newText.isNullOrEmpty()) {
                    return true
                }

                val spannableString = SpannableString(contentText)


                val startIndex = contentText.indexOf(newText, ignoreCase = true)
                val endIndex = startIndex + newText.length


                if (startIndex != -1) {
                    spannableString.setSpan(
                        ForegroundColorSpan(resources.getColor(R.color.purple_700)),
                        startIndex,
                        endIndex,
                        Spannable.SPAN_EXCLUSIVE_EXCLUSIVE
                    )
                }


                chapterContent.text = spannableString

                return true
            }
        })
    }

}