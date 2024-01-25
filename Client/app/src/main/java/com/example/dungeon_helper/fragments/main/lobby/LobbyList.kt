package com.example.dungeon_helper.fragments.main.lobby

import androidx.lifecycle.ViewModelProvider
import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import com.example.dungeon_helper.MainActivity
import com.example.dungeon_helper.R
import com.example.dungeon_helper.databinding.FragmentLobbyListBinding

class LobbyList : Fragment() {

    companion object {
        fun newInstance() = LobbyList()
    }

    private lateinit var viewModel: LobbyListViewModel

    private var _binding: FragmentLobbyListBinding? = null
    private  val binding get() = _binding!!

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        val lobbyMainViewModel = ViewModelProvider(this)[LobbyListViewModel::class.java]

        _binding = FragmentLobbyListBinding.inflate(inflater, container, false)
        val root: View = binding.root

        val textView: TextView = binding.textLobbyList
        lobbyMainViewModel.text.observe(viewLifecycleOwner) {
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

        val backBtn = binding.backBtn


        backBtn.setOnClickListener {
            (activity as MainActivity).navController.navigate(R.id.action_lobbyList_to_navigation_lobby)
        }

    }

}