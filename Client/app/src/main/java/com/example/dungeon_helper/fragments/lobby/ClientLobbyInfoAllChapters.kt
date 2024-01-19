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
import com.example.dungeon_helper.ClientLobbyActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentClientLobbyInfoAllChaptersBinding

class ClientLobbyInfoAllChapters : Fragment() {

    companion object {
        fun newInstance() = ClientLobbyInfoAllChapters()
    }

    private lateinit var viewModel: ClientLobbyInfoAllChaptersViewModel
    private var _binding: FragmentClientLobbyInfoAllChaptersBinding? = null
    private val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
       val clientLobbyInfoAllChaptersViewModel = ViewModelProvider(this)[ClientLobbyInfoAllChaptersViewModel::class.java]
        _binding = FragmentClientLobbyInfoAllChaptersBinding.inflate(inflater, container, false)
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
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobbyInfoAllChapters_to_clientLobby)
        }
        val gameChaptersView = binding.gameChapter1
        gameChaptersView.setOnClickListener{
            (activity as ClientLobbyActivity).navController.navigate(R.id.action_clientLobbyInfoAllChapters_to_clientLobbyInfoContent)
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

                val gameChaptersText = gameChaptersView.text.toString()
                val spannableString = SpannableString(gameChaptersText)


                val startIndex = gameChaptersText.indexOf(newText, ignoreCase = true)
                val endIndex = startIndex + newText.length

                if (startIndex != -1) {
                    spannableString.setSpan(
                        ForegroundColorSpan(resources.getColor(R.color.purple_700)),
                        startIndex,
                        endIndex,
                        Spannable.SPAN_EXCLUSIVE_EXCLUSIVE
                    )
                }


                gameChaptersView.text = spannableString

                return true
            }
        })

    }

}